import React, { Component } from 'react';
import styled from 'styled-components';

import MapContainer from './MapContainer';
import Selected from './Selected';

class App extends Component {
  render() {
    return (
      <Container>
        <Selected />
        <MapContainer/>
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

