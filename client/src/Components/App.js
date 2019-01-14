import React, { Component } from 'react';
import styled from 'styled-components';

import { Route } from 'react-router-dom';

import MapContainer from './MapContainer';
import Header from './Header';
import Register from './Register';
import Login from './Login';

class App extends Component {
  render() {
    return (
      <Container>
        <Header/>
        <MapContainer/>
        <Route path="/login" component={props => <Login {...props}/>}/>
        <Route path="/register" component={props => <Register {...props}/>}/>
      </Container>
    );
  }
}

export default App;

const Container = styled.div`
  width: 100%;
  max-width: 900px;
  min-width: 600px;
  margin: 0 auto;
  height: 100vh;
  padding: 50px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  position: relative;
  /* border: solid 1px red; */
`;

