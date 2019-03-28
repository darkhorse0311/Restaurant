import React, { Component } from 'react';
import styled from 'styled-components';

import { Route } from 'react-router-dom';

import Map from './Map/container';
import Items from './Items/container';
import Navigation from './Navigation/container';

class App extends Component {
  render() {
    return (
      <Container>
        <Navigation/>
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
  background-color: #2B2B2B;
  /* border: solid 1px red; */
`;

