import React from 'react';
import styled from 'styled-components'

const Arrows = () => {
    return (
        <RightArrow>
            <i className="fas fa-angle-right"></i>
        </RightArrow>
    );
}

export default Arrows;

const RightArrow = styled.span`
    position: absolute;
    right: 0;
    top: 0;
    height: 100%;
    width: 100px;
    border-top-right-radius: 12px;
    z-index: 2;
    background-color: white;
    color: black;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 40px;
`;
