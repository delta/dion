import { useContext } from "react";
import { Link, Navigate } from "react-router-dom";

import deltaLogo from "../../assets/deltaLogoWhite.png";
import { AuthContext } from "../../context/user-context";

const DAuthLogin = () => {
  const context = useContext(AuthContext);
  if (context == null) {
    return <h1>Loading</h1>;
  }

  if (context.email != null) {
    return <Navigate replace to="/game" />;
  }

  return (
    <div>
      <span>
        <Link to="/auth/callback">
          <button>
            Login With
            <img src={deltaLogo} alt="your mom" />
          </button>
        </Link>
      </span>
    </div>
  );
};

export default DAuthLogin;
