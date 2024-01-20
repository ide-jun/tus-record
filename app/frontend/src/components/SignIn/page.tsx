import axios from "axios";
import { Footer, Header } from "../Utils/CommonForms";
import SignInForm from "./form";
import { useNavigate } from "react-router-dom";

export const SignInPage = () => {
    const navigate = useNavigate();
    const handleSignUp = (email: string, password: string) => {
        axios.post('http://localhost:8080/api/auth/sign-in', {
            "email": email, 
            "password": password
        }).then(res => {
            localStorage.setItem("token", res.data["token"])
            navigate("/post-record")
        });
    };

    return (
        <div className="App">
            <Header />
            <SignInForm onSubmit={handleSignUp} />
            <Footer />
        </div>
    );
}