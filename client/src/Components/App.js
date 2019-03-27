import React, { Component } from 'react';
import styled from 'styled-components';

import { Route } from 'react-router-dom';

import Map from './Map/container';
import Header from './Header';
import Items from './Items';

class App extends Component {
  render() {
    return (
      <Container>
        <Header/>
        <Route path="/" render={props => <Map {...props}/>} />
        <Route path="/items" render={props => <Items {...props}/>}/>
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

