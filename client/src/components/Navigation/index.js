import React from 'react';
import styled, { keyframes } from 'styled-components'

const Navigation = ({center, getLocations, loading}) => {
    return (
        <StyledHeader>
            <h1>Carbtographer</h1>
            {
                loading
                ? (<i className="fas fa-spinner spinner"></i>)
                : (
                <i
                    className="fas fa-redo-alt"
                    onClick={() => {
                        getLocations(center[0], center[1])
                    }}
                />
                )
            }
        </StyledHeader>
    );
}

export default Navigation;

const rotation = keyframes`
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(359deg);
    }
`

const StyledHeader = styled.div`
    width: 100vw;
    height: 60px;
    position: fixed;
    z-index: 2;
    top: 0;
    left: 0;
    display: flex;
    padding: 0 20px;
    justify-content: space-between;
    align-items: center;
    color: rgba(255,255,255,0.74);
    i {
        font-size: 26px;
        height: 60px;
        width: 60px;
        line-height: 60px;
        text-align: center;
        &.spinner {
            animation: ${rotation} 2.5s infinite linear;
        }
    }
    h1 {
        line-height: 60px;
        margin: 0;
    }
`;

