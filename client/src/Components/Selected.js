import React, { Component } from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components'


class Selected extends Component {
    render() {
        if (this.props.selected === null) {
            return (
                <Container>
                    <h3>Select Restuarant</h3>
                </Container>
            )
        }
        const { name } = this.props.selected;
        return (
            <Container>
                <h3>{name}</h3>
            </Container>
        );
    }
}

const mapStateToProps = state => {
    const { selected } = state.info
    return {
        selected
    }
}

export default connect(mapStateToProps, {})(Selected);

const Container = styled.div`
    width: 200px;
    height: 350px;
    border: 1px solid rgba(0,0,0,0.54);
    border-radius: 8px;
    position: absolute;
    left: 25px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;

`;
