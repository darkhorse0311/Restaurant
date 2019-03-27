import React from 'react';
import styled from 'styled-components'
import Item from './Item';
import Arrows from './Arrows';


// {place && (
//     <BottomPopup place={place} clearPlace={this.clearPlace} />
//   )}

const BottomPopup = ({place, clearPlace}) => {
    return (
        <StyledContainer>
            <Header>
                <h2>{place.name}</h2>
                <span onClick={clearPlace}>X</span>
            </Header>
            <ItemList>
            {
                place.items.map((item, i) => (
                    <Item key={i} item={item}/>
                ))
            }
            </ItemList>
            <Arrows/>
        </StyledContainer>
    );
}

export default BottomPopup;

const ItemList = styled.div`
    height: 165px;
    display: flex;
    justify-content: flex-start;
    padding-bottom: 20px;
    position: relative;
    left: -100%;
    transform: translateX(100%);
`;

const Header = styled.div`
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 54px;
    width: 100%;
    span {
        z-index: 3;
    }
`;

const StyledContainer = styled.div`
    position: fixed;
    bottom: 0;
    width: 100vw;
    height: 254px;
    background: white;
    z-index: 99;
    padding: 0px 40px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    overflow: hidden;
`;

