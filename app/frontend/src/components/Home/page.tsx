import { assertExpressionStatement } from "@babel/types";
import { Link } from "react-router-dom";
import axios from 'axios'
import Ping from "./ping-form";
import { useState } from "react";
import { Header, Footer } from "../Utils/Common";

export const Home = () => {
  const token = localStorage.getItem("token")
  const handlePing = () => {
    axios
      .post('http://localhost:8080/api/ping', {}, {headers: {"Authorization": `Bearer ${token}`}})
      .then(any => {
        console.log(any.data["receivedTime"])
      });
  };
  
  return (
  <div className='App'>
    <Header />
    <Ping onSubmit={handlePing} />
    <li>
      <Link to="/sign-in">サインイン</Link>
    </li>
    <li>
      <Link to="/sign-up">サインアップ</Link>
    </li>
    <Footer />
  </div>
  );
}