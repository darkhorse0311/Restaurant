import React from 'react';
import styled, { css } from 'styled-components'
import Item from './Item';


const ItemsContainer = ({name, items, setShowModal, showModal, setSortMode, sortMode, order}) => {

    const changeMode = (mode) => {
        let nOrder = 'A';
        if (sortMode === mode) {
            nOrder = (order === 'A') ? 'D' : 'A';
        }

        setSortMode(mode, nOrder);
    }

    const sortItems = (items) => {
        let key = sortMode === 'C'
            ? 'carbs'
                : sortMode === 'P'
                ? 'protein'
                    : sortMode === 'F'
                    ? 'fats'
                        : null;

        if (key == null) {
            return items
        }

        let sItems = items.sort((a, b) => {
            return a[key] - b[key]
        });

        if (order === 'D') {
            return sItems.reverse();
        }

        return sItems;
    }

    return (
        <StyledContainer showModal={showModal}>
            <Header>
                <div className={"title"}>
                    <h2>{name}</h2>
                    <i 
                        onClick={() => setShowModal(false)}
                        className={"down fas fa-times"}
                    />
                </div>
                <div className={"actions"}>
                    <span onClick={() => changeMode('C')}>CARBS</span>
                    <span onClick={() => changeMode('P')}>PROTEIN</span>
                    <span onClick={() => changeMode('F')}>FATS</span>
                </div>
            </Header>
            <ItemList>
            {
                items.length ? sortItems(items).map((item, i) => (
                    <Item key={i} item={item}/>
                )) : null
            }
            </ItemList>
        </StyledContainer>
    );
}

export default ItemsContainer;



const Header = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
    padding: 10px 0px 10px;
    .actions {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 60px;
        font-size: 15px;
        font-weight: 600;
        text-transform: uppercase;
        cursor: pointer;
    }
    .title {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 40px;
    }
    .down {
        font-size: 26px;
    }
    h2 {
        font-size: 26px;
        font-weight: 600;
    }
`;

const ItemList = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    position: relative;
    border-top: 1px solid black;
    overflow: scroll;
`;

const StyledContainer = styled.div`
    width: 95%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    height: calc(100% - 60px);
    bottom: 0;
    right: 2.5%;
    padding: 10px 20px 0px;
    position: absolute;
    background: white;
    z-index: 1;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    transform: translateY(100%);
    transition: transform 300ms;
    ${props => props.showModal && css`
        transform: translateY(0);
    `}
    ${props => !props.showModal && css`
        transform: translateY(100%);
    `}

    @media (min-width: 650px) {
        width: 325px;
        position: absolute;
        height: 100vh;
        right: 0;
        margin-left: 0px;
        transition: transform none;
        transform: translateY(0);
        border-top-right-radius: 0px;
    }
`;

