import React from 'react';
import styled from 'styled-components'

const Header = () => {
    return (
        <StyledHeader>
            <h1>Carbtographer</h1>
            <StyledButton>Signup</StyledButton>
        </StyledHeader>
    );
}

export default Header;

const StyledHeader = styled.div`
    width: 100vw;
    height: 60px;
    position: fixed;
    z-index: 2;
    top: 0;
    left: 0;
    display: flex;
    padding: 0 30px;
    justify-content: space-between;
    align-items: center;
    color: rgba(255,255,255,0.74);
`;

const StyledButton = styled.span`
    width: 140px;
    height: 36px;
    font-size: 18px;
    line-height: 32px;
    border: 2px solid rgba(255,255,255,0.74);
    text-align: center;
    border-radius: 18px;
    font-weight: 600;
    cursor: pointer;
`;
