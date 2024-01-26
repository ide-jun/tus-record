import { useState } from "react"
import '../SignIn/form.css'

interface PingProps {
    onSubmit: () => void
}

const Ping: React.FC<PingProps> = ({ onSubmit }) => {

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        onSubmit();
      };

    return (
        <form className="form" onSubmit={handleSubmit}>
            <h1>疎通確認</h1>
            <button type="submit">Ping</button>
        </form>
    );
};

export default Ping