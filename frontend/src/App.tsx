import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
// import init, {encrypt, new_key} from "encrypt";

function App() {
  const [count, setCount] = useState(0)
  // Dummy code that works
  // useEffect(() => {
  //   init().then(() => {
  //     const a = encrypt(JSON.stringify({"a": "b"}), new_key());
  //     console.log(a.get_nonce());
  //     console.log(a.get_value());
  //   })
  // })

  return (
		<>
			<h1 className="text-3xl font-bold">
				Welcome to <span className="text-blue-400">Dion</span>
			</h1>
			;
		</>
	);
}

export default App
