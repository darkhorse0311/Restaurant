import React from 'react';
import styled, { css } from 'styled-components'
import Item from './Item';


const ItemsContainer = ({name, items, setShowModal, showModal}) => {
    return (
        <StyledContainer showModal={showModal}>
            <Header>
                <h2>{name}</h2>
                <i 
                    onClick={() => setShowModal(false)}
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


const StyledContainer = styled.div`
    width: 95%;
    height: calc(100% - 60px);
    bottom: 0;
    left: 2.5%;
    padding: 10px 20px 0px;
    position: absolute;
    background: white;
    z-index: 1;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    overflow: scroll;
    transform: translateY(100%);
    transition: transform 300ms;
    ${props => props.showModal && css`
        transform: translateY(0);
    `}
    ${props => !props.showModal && css`
        transform: translateY(100%);
    `}
`;

const ItemList = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    position: relative;
    border-top: 1px solid black;
    margin-top: 74px;
`;

const Header = styled.div`
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 74px;
    width: 100%;
    padding: 10px 30px 10px;
    position: absolute;
    z-index: 2;
    left: 0;
    top: 0;
    .down {
        font-size: 26px;
    }
    h2 {
        font-size: 26px;
        font-weight: 600;
    }
`;

