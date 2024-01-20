import axios from "axios";
import { Footer, Header } from "../Utils/CommonForms";
import SignUpForm from "./form";

export const SignUpPage = () => {
    const handleSignUp = (name: string, email: string, password: string) => {
        axios.post('http://localhost:8080/api/auth/sign-up', {
            "name": name, 
            "email": email, 
            "password": password
        });
    };

    return (
        <div className="App">
            <Header />
            <SignUpForm onSubmit={handleSignUp} />
            <Footer />
        </div>
    );
}