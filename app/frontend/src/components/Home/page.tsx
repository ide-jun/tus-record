import { assertExpressionStatement } from "@babel/types";
import { Link } from "react-router-dom";
import axios from 'axios'
import Ping from "./form";

export const Home = () => {
  const handlePing = (message: string) => {
    axios.post('http://localhost:8080/ping', {message: message});
  };
  
  return (
  <div className='App'>
    <h1>
      tus-record
    </h1>
    <Ping onSubmit={handlePing} />
    <li>
      <Link to="/sign-in">サインイン</Link>
    </li>
    <li>
      <Link to="/sign-up">サインアップ</Link>
    </li>
  </div>
  );
}