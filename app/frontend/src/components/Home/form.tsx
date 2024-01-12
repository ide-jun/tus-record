import { useState } from "react"

interface PingProps {
    onSubmit: (message: string) => void
}

const Ping: React.FC<PingProps> = ({ onSubmit }) => {
    const [message, setMessage] = useState('')

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        onSubmit(message);
      };

    return (
        <form onSubmit={handleSubmit}>
            <h1>疎通確認</h1>
            <label>
                メッセージ
                <input type="name" value={message} onChange={(e) => setMessage(e.target.value)} />
            </label>
            <button type="submit">Ping</button>
        </form>
    );
};

export default Ping