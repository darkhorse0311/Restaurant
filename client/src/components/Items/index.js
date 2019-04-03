import React from 'react';
import styled, { css } from 'styled-components'
import Item from './Item';


const ItemsContainer = ({name, items, setShowModal, showModal, setSortMode, sortMode, order}) => {

    const changeMode = (mode) => {
        let nOrder = 'A';

        if (order === 'D'){
            setSortMode('N', 'A');
        } else {
            if (sortMode === mode) {
                nOrder = (order === 'A') ? 'D' : 'A';
            }
            setSortMode(mode, nOrder);
        }

    }

    const sortItems = () => {

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

        let sItems = [...items].sort((a, b) => {
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
                        className={"close fas fa-times"}
                    />
                </div>
                <ActionGroup>
                    <StyledAction 
                        mode={sortMode}
                        order={order}
                        type={'C'}
                        onClick={() => changeMode('C')}
                    >
                        <span>CARBS</span>
                        <span className="arrow-box">
                            <i className="fas fa-sort-up"></i>
                            <i className="fas fa-sort-down"></i>
                        </span>
                    </StyledAction>
                    <StyledAction 
                        onClick={() => changeMode('P')}
                        mode={sortMode}
                        order={order}
                        type={'P'}
                    >
                        <span>PROTEIN</span>
                        <span className="arrow-box">
                            <i className="fas fa-sort-up"></i>
                            <i className="fas fa-sort-down"></i>
                        </span>
                    </StyledAction>
                    <StyledAction 
                        onClick={() => changeMode('F')}
                        mode={sortMode}
                        order={order}
                        type={'F'}
                    >
                        <span>FATS</span>
                        <span className="arrow-box">
                            <i className="fas fa-sort-up"></i>
                            <i className="fas fa-sort-down"></i>
                        </span>
                    </StyledAction>
                </ActionGroup>
            </Header>
            <ItemList>
            {
                items.length ? sortItems().map((item, i) => (
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
    .title {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 40px;
    }
    .close {
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

const StyledAction = styled.div`
    display: flex;
    justify-content: flex-start;
    align-items: center;
    user-select: none;
    .arrow-box {
        height: 30px;
        font-size: 14px;
        position: relative;
        margin: 0px 5px;
        .fa-sort-up {
            position: absolute;
            top:8px;
            opacity: 0.3;
            ${({mode, order, type}) => (mode === type && order === 'A') && css`
                opacity: 1;
            `}
        }
        .fa-sort-down {
            position: absolute;
            top:10px;
            opacity: 0.3;
            ${({mode, order, type}) => (mode === type && order === 'D') && css`
                opacity: 1;
            `}
        }
    }
`;

const ActionGroup = styled.div`
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 60px;
    font-size: 15px;
    font-weight: 600;
    text-transform: uppercase;
    cursor: pointer;
`
