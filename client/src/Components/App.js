import React, { Component } from 'react';
import styled from 'styled-components';

import { Route } from 'react-router-dom';

import Map from './Map/container';
import Items from './Items/container';
import Header from './shared/Header';

class App extends Component {
  render() {
    return (
      <Container>
        <Header/>
        <Route path="/" render={props => (
          <>
            <Map {...props}/>
            <Items {...props}/>
          </>
        )} />
      </Container>
    );
  }
}

export default App;

const Container = styled.div`
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  position: relative;
  /* border: solid 1px red; */
`;

