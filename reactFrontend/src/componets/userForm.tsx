import React, { useState } from 'react';
import axios from 'axios';

function UserForm() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post('/users/create', { name, email, password })
      .then(response => {
        console.log(response);
      })
      .catch(error => {
        console.log(error);
      });
  };

  return (
    <div>
      <h2>Add User</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input type="text" value={name} onChange={e => setName(e.target.value)} />
        </label>
        <label>
          Email:
          <input type="email" value={email} onChange={e => setEmail (e.target.value)} />
        </label>
       <label>
          Password:
         <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
        </label>
       <button type="submit">Submit</button>
      </form>
    </div>
  );
}

export default UserForm;