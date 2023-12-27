import React, { useState, useEffect } from 'react';
import { Routes, Route } from 'react-router-dom'
import { Home } from './components/Home'
import { About } from './components/About'
import { Contact } from './components/Contact'
import { NotFound } from './components/NotFound';
import { Footer } from './components/Footer'
import SignIn from './components/signin/sign-in-form'

import axios from "axios";

function App() {
  const handleSignin = (name: string, email: string, password: string) => {
    axios.post('http://localhost:8080/sign-in', {name: name, email: email, password: password});
  };

  return (
    <div className='App'>
      <h1>
        tus-record
      </h1>
      <SignIn onSubmit={handleSignin}/>
      <Routes>
        <Route path="/" element={ <Home/> }/>
        <Route path="/about" element={ <About/> }/>
        <Route path="/contact" element={ <Contact/> }/>
        <Route path="*" element={ <NotFound /> } />
      </Routes>
      <Footer />
    </div>
  )
  // const [fruits, setFruits] = useState<Fruit[]>([{ id: 0, name: "", icon: "" }])

  // useEffect(() => {
  //   (
  //     async () => {
  //       const data = await axios.get("http://localhost:8080/fruit")
  //       console.log(data.data)
  //       console.log(data.data[0])
  //       setFruits(data.data)
  //     }
  //   )()
  // }, [])

  // return (
  //   <div>
  //     {fruits.map(fruit => (
  //       <p key={fruit.id}>
  //         <span>{fruit.name}</span>
  //         <span>{fruit.icon}</span>
  //       </p>
  //     )
  //     )
  //     }
  //   </div>
  // );
}

export default App;

