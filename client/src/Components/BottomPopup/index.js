import React from 'react';
import styled from 'styled-components'

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
                    <Item key={i}>
                        <h3>{item.name}</h3>
                        <ItemMacros>
                            <Macro>
                                <Gram>
                                    {item.carbs}
                                </Gram>
                                <MacroTitle>
                                    Carbs
                                </MacroTitle>
                            </Macro>
                            <Macro>
                                <Gram>
                                    {item.protein}
                                </Gram>
                                <MacroTitle>
                                    Protein
                                </MacroTitle>
                            </Macro>
                            <Macro>
                                <Gram>
                                    {item.fats}
                                </Gram>
                                <MacroTitle>
                                    Fats
                                </MacroTitle>
                            </Macro>
                        </ItemMacros>
                    </Item>
                ))
            }
            </ItemList>
        </StyledContainer>
    );
}

export default BottomPopup;

const MacroTitle = styled.span`
    font-size: 12px;
    font-weight: 600;
    text-transform: uppercase;
    text-align: center;
`;

const Gram = styled.span`
    font-size: 22px;
    font-weight: 800;
`;

const Macro = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`;

const ItemMacros = styled.div`
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    width: 100%;
`;

const Item = styled.div`
    width: 200px;
    min-width: 200px;
    height: 100%;
    border: 1px solid rgba(0,0,0,0.4);
    border-radius: 12px;
    margin-right: 10px;
    position: relative;
    padding: 10px 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: flex-start;

    h3 {
        margin: 0px;
    }
`;

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

