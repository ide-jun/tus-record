import { Route, Routes } from "react-router"
import { Home } from "./components/Home/page"
import { SignIn } from "./components/SignIn/page"
import { SignUp } from "./components/SignUp/page"

export const Router = () => {
    return (
    <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/sign-in" element={<SignIn />} />
        <Route path="/sign-up" element={<SignUp />} />
    </Routes>
    )
}