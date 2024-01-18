import axios from "axios";
import { Footer, Header } from "../Utils/Common";
import SignInForm from "./form";

export const SignInPage = () => {
    const handleSignUp = (email: string, password: string) => {
        axios.post('http://localhost:8080/api/auth/sign-in', {
            "email": email, 
            "password": password
        }).then(res => {
            localStorage.setItem("token", res.data["token"])
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