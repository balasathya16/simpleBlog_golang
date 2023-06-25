import React from 'react';
import './Header.css';

const Header = () => {
  return (
    <header className="header">
  <h1 className="title">A Failed Cricketer</h1>
  <div className="spacer"></div>
  <div className="button-container">
    <ul className="nav-list">
      <li className="nav-item"><a href="/blogs">Blogs</a></li>
    </ul>
    <button className="login-button">Login</button>
    <button className="start-button">Get Started</button>
  </div>
  <div className="spacer"></div>
</header>
  
  );
};

export default Header;
