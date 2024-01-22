import React, { useState } from 'react';
// import './form.css'

interface SignInProps {
  onSubmit: (email: string, password: string) => void;
}

const PostRecordForm: React.FC<SignInProps> = ({ onSubmit }) => {
  const [email,    setEmail]    = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(email, password);
  };

  return (
    <form className="form" onSubmit={handleSubmit}>
      <h1>SignIn</h1>
      <label>
        メールアドレス
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      </label>
      <label>
        パスワード
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </label>
      <button type="submit">SignIn</button>
    </form>
  );
};

export default PostRecordForm;
