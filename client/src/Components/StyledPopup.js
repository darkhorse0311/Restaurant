import React from 'react';
import styled from 'styled-components';

const StyledPopup = ({place, clearPlace}) => {
    return (
        <StyledPopupContainer>
            <PopupHeader>   
                {place.name}
                <span onClick={() => clearPlace()}>x</span>
            </PopupHeader>
            <ListWrapper>
            {
                place.items.map((item, i) => (
                    <div key={i}>
                        <h4>{item.name}</h4>
                        <div>Protein: {item.protein}</div>
                        <div>Carbs: {item.carbs}</div>
                        <div>Fats: {item.fats}</div>
                    </div>
                ))
            }
            </ListWrapper>
        </StyledPopupContainer>
    );
}

export default StyledPopup;

const StyledPopupContainer = styled.div`
  background: white;
  color: #3f618c;
  font-weight: 400;
  padding: 5px;
  border-radius: 2px;
  max-height: 200px;
  overflow: scroll;
`;

const ListWrapper = styled.div`
    width: 100%;
    overflow: scroll;
    margin-top: 10px;
`;

const PopupHeader = styled.div`
    position: absolute;
    top: 10px;
    left: 0px;
    padding: 5px 10px;
    width: 100%;
    background-color: white;
    span {
        position: absolute;
        right: 5px;
        top: 0px;
        font-weight: 800;
        width: 20px;
        height: 20px;
        user-select: none;
        cursor: pointer;
    }
`;
