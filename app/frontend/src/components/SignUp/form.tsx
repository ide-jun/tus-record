import React, { useState } from 'react';
import '../SignIn/form.css'

interface SignUpProps {
  onSubmit: (name: string, email: string, password: string) => void;
}

const SignUpForm: React.FC<SignUpProps> = ({ onSubmit }) => {
  const [name,     setName]     = useState('')
  const [email,    setEmail]    = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(name, email, password);
  };

  return (
    <form className="form" onSubmit={handleSubmit}>
      <h1>SignUp</h1>
      <label>
        ユーザー名
        <input type="name" value={name} onChange={(e) => setName(e.target.value)} />
      </label>
      <label>
        メールアドレス
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      </label>
      <label>
        パスワード
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </label>
      <button type="submit">SignUp</button>
    </form>
  );
};

export default SignUpForm;