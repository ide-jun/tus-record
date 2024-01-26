import { assertExpressionStatement } from "@babel/types";
import axios, { HttpStatusCode } from 'axios'
import { useEffect, useState } from "react";
import { Header, Footer, AuthErrorForm } from "../Utils/CommonForms";
import { error } from "console";
import PostRecordForm from "./form";

export const PostRecordPage = () => {
    const [name, setName] = useState("")
    const [auth, setAuth] = useState(false)
  const token = localStorage.getItem("token")
  useEffect(() => {
    const getData = () => {
        axios
            .get("http://localhost:8080/api/auth/get-data", {headers: {"Authorization": `Bearer ${token}`}})
            .then(res => {
                setName(res.data["name"])
                setAuth(true)
            })
            .catch(err => {
                if (axios.isAxiosError(err) && err.response?.status === HttpStatusCode.Unauthorized) {
                    setAuth(false)
                }
            })
    }
    getData()
  }, [])
  const handlePing = () => {
    axios
      .post('http://localhost:8080/api/post-record/post', {}, {headers: {"Authorization": `Bearer ${token}`}})
      .then(any => {
        console.log(any.data["receivedTime"])
      });
  };
  
    if (auth) {
  return (
  <div className='App'>
    <Header />
    <p> ログイン名：{name}</p>
    <PostRecordForm onSubmit={handlePing} />
    <Footer />
  </div>
  );
    }
    return (
        <div className='App'>
        <Header />
            <AuthErrorForm />
        <Footer />
    </div>
    )
}