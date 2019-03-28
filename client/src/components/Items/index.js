import React from 'react';
import styled from 'styled-components'
import Item from './Item';


const ItemsContainer = ({name, items, history}) => {
    return (
        <StyledContainer>
            <Header>
                <h2>{name}</h2>
                <i 
                    onClick={() => history.push('/')}
                    className={"down fas fa-long-arrow-alt-down"}
                />
            </Header>
            <ItemList>
            {
                items.length ? items.map((item, i) => (
                    <Item key={i} item={item}/>
                )) : null
            }
            </ItemList>
        </StyledContainer>
    );
}

export default ItemsContainer;

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
    padding: 0px 10px;
    .down {
        font-size: 26px;
    }
    h2 {
        font-size: 26px;
        font-weight: 600;
    }
`;

const StyledContainer = styled.div`
    width: 95%;
    height: 95%;
    bottom: 0;
    left: 2.5%;
    padding: 10px 10px 0px;
    position: absolute;
    background: white;
    z-index: 999999;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    overflow: hidden;
`;

