import { assertExpressionStatement } from "@babel/types";
import { Link } from "react-router-dom";
import axios from 'axios'
import Ping from "./ping-form";
import { useState } from "react";
import { Header, Footer } from "../Utils/Common";

export const Home = () => {
  const [msg, setMsg] = useState('')
  const [time, setTime] = useState('')
  const handlePing = (message: string) => {
    axios
      .post('http://localhost:8080/ping', {message: message})
      .then(response => {
        setMsg(response.data["message"]);
        setTime(response.data["receivedTime"]);
      });
  };
  
  return (
  <div className='App'>
    <Header />
    <Ping onSubmit={handlePing} />
    <h1>メッセージ</h1>{msg}
    <h1>受信時間</h1>{time}
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