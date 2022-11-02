mod utils;

use aes_gcm::{
    aead::{generic_array::GenericArray, Aead, KeyInit, OsRng},
    Aes256Gcm,
    Nonce, // Or `Aes128Gcm`
};
use base64ct::{Base64, Encoding};
use hex;
use rand::Rng;
use serde::{Deserialize, Serialize};
use serde_json::Value;
use sha2::{Digest, Sha256};
use std::collections::HashMap;
use std::convert::TryInto;
use std::panic;
use wasm_bindgen::prelude::*;

// When the `wee_alloc` feature is enabled, use `wee_alloc` as the global
// allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

#[derive(Debug, Serialize, Deserialize)]
struct EncryptedValues {
    hashed_key: String,
    value: String,
}

#[wasm_bindgen]
pub struct EncryptReturn {
    nonce: String,
    value: String,
}

#[wasm_bindgen]
impl EncryptReturn {
    pub fn get_nonce(&self) -> String {
        self.nonce.clone()
    }

    pub fn get_value(&self) -> String {
        self.value.clone()
    }
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(untagged)]
enum HashedValue {
    Null,
    Value(EncryptedValues),
    Object(HashMap<String, HashedValue>),
}

enum CurrentKey {
    Null,
    Value(String),
}

impl EncryptedValues {
    pub fn new(hashed_key: String, value: String) -> Self {
        EncryptedValues { hashed_key, value }
    }
}

struct Encryptor {
    aes: Aes256Gcm,
    nonce: [u8; 12],
}

impl Encryptor {
    pub fn new(key: String, nonce: [u8; 12]) -> Self {
        Self {
            aes: Aes256Gcm::new(GenericArray::from_slice(
                hex::decode(&key).unwrap().as_slice(),
            )),
            nonce,
        }
    }

    fn hash_encrypt_value(
        &self,
        key: &CurrentKey,
        value: &String,
    ) -> Result<(String, String), String> {
        let nonce = Nonce::from_slice(&self.nonce);
        let key_string = match key {
            CurrentKey::Null => "",
            CurrentKey::Value(x) => x,
        };
        Ok((
            Base64::encode_string(&Sha256::digest(key_string)),
            Base64::encode_string(
                &self
                    .aes
                    .encrypt(nonce, value.as_bytes())
                    .map_err(|e| e.to_string())?,
            ),
        ))
    }

    pub fn hash_encrypt_object(
        &self,
        to_encrypt: &Value,
        key: &CurrentKey,
    ) -> Result<HashedValue, String> {
        match to_encrypt {
            Value::Object(obj) => {
                let nonce = Nonce::from_slice(&self.nonce);
                let mut res = HashMap::new();
                for (key, val) in obj {
                    let encrypted_key = Base64::encode_string(
                        &self
                            .aes
                            .encrypt(nonce, key.as_bytes())
                            .map_err(|e| e.to_string())?,
                    );
                    res.insert(
                        encrypted_key.clone(),
                        self.hash_encrypt_object(val, &CurrentKey::Value(key.clone()))?,
                    );
                }
                Ok(HashedValue::Object(res))
            }
            Value::String(s) => {
                let (key, value) = self.hash_encrypt_value(&key, s)?;
                Ok(HashedValue::Value(EncryptedValues::new(key, value)))
            }
            Value::Bool(b) => {
                let (key, value) = self.hash_encrypt_value(&key, &b.to_string())?;
                Ok(HashedValue::Value(EncryptedValues::new(key, value)))
            }
            Value::Number(n) => {
                let (key, value) = self.hash_encrypt_value(&key, &n.to_string())?;
                Ok(HashedValue::Value(EncryptedValues::new(key, value)))
            }
            Value::Null => Ok(HashedValue::Null),
            Value::Array(a) => {
                let v = serde_json::to_string(a).map_err(|e| e.to_string())?;
                let (key, value) = self.hash_encrypt_value(&key, &v.to_string())?;
                Ok(HashedValue::Value(EncryptedValues::new(key, value)))
            }
        }
    }

    pub fn decrypt_object(&self, to_decrypt: HashedValue) -> Result<Value, String> {
        let nonce = Nonce::from_slice(&self.nonce);
        match to_decrypt {
            HashedValue::Null => Ok(Value::Null),
            HashedValue::Value(EncryptedValues {
                hashed_key: _,
                value,
            }) => {
                let new_value = Base64::decode_vec(&value).unwrap();
                Ok(Value::String(
                    String::from_utf8(
                        self.aes
                            .decrypt(nonce, new_value.as_slice())
                            .map_err(|e| e.to_string())?,
                    )
                    .map_err(|e| e.to_string())?,
                ))
            }
            HashedValue::Object(x) => {
                let mut res = serde_json::Map::new();
                for (key, val) in x {
                    let m = self
                        .aes
                        .decrypt(
                            nonce,
                            Base64::decode_vec(&key)
                                .map_err(|e| e.to_string())?
                                .as_slice(),
                        )
                        .map_err(|e| e.to_string())?;
                    res.insert(String::from_utf8(m).unwrap(), self.decrypt_object(val)?);
                }
                Ok(Value::Object(res))
            }
        }
    }
}

fn new_nonce() -> [u8; 12] {
    rand::thread_rng().gen::<[u8; 12]>()
}

#[wasm_bindgen]
pub fn new_key() -> String {
    hex::encode(Aes256Gcm::generate_key(&mut OsRng))
}

#[wasm_bindgen]
pub fn encrypt(data: String, key: String) -> Result<EncryptReturn, String> {
    utils::set_panic_hook();
    let v: Value = serde_json::from_str(&data).unwrap();
    let nonce = new_nonce();
    let encryptor = Encryptor::new(key, nonce);
    let res = encryptor.hash_encrypt_object(&v, &CurrentKey::Null)?;
    match serde_json::to_string(&res) {
        Ok(r) => Ok(EncryptReturn {
            nonce: hex::encode(nonce),
            value: r,
        }),
        Err(_) => Err("Couldn't parse".to_string()),
    }
}

#[wasm_bindgen]
pub fn decrypt(data: String, key: String, nonce: String) -> Result<String, String> {
    panic::set_hook(Box::new(console_error_panic_hook::hook));
    let v: HashedValue = serde_json::from_str(&data).unwrap();
    let nonce = hex::decode(&nonce).unwrap();
    let nonce: [u8; 12] = match nonce.try_into() {
        Ok(n) => n,
        Err(_) => {
            return Err("Invalid nonce".to_string());
        }
    };
    let decryptor = Encryptor::new(key, nonce);
    let res = decryptor.decrypt_object(v)?;
    serde_json::to_string(&res).map_err(|_e| "Couldn't parse".to_string())
}

