import { Route, Routes } from "react-router"
import { Home } from "./components/Home/page"
import { SignInPage } from "./components/SignIn/page"
import { SignUpPage } from "./components/SignUp/page"
import { PostRecordPage } from "./components/PostRecord/page"

export const Router = () => {
    return (
    <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/sign-in" element={<SignInPage />} />
        <Route path="/sign-up" element={<SignUpPage />} />
        <Route path="/post-record" element={<PostRecordPage />} />
    </Routes>
    )
}