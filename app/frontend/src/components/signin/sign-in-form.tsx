import React, { useState } from 'react';
import styles from './sign-in-form.module.scss'

interface SignInProps {
  onSubmit: (name: string, email: string, password: string) => void;
}

const SignIn: React.FC<SignInProps> = ({ onSubmit }) => {
  const [name,     setName]     = useState('')
  const [email,    setEmail]    = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(name, email, password);
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <h1>SignIn</h1>
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
      <button type="submit">SignIn</button>
    </form>
  );
};

export default SignIn;
