import React from 'react';
import './Home.scss';

export default class Home extends React.Component {
  render() {
    console.log("got to home");
    return (
      <div className='home-container'>
        <b>hi there!</b>
      </div>
    );
  }
}
