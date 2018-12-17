import React, { Component } from 'react';
import styled from 'styled-components';

import MapContainer from './MapContainer';

class App extends Component {
  render() {
    return (
      <Container>
        <MapContainer/>
      </Container>
    );
  }
}

export default App;

const Container = styled.div`
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
`;

