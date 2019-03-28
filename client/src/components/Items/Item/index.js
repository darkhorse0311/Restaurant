import React from 'react';
import styled from 'styled-components';

const Item = ({item}) => {
    return (
        <StyledItem>
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
        </StyledItem>
    );
}

export default Item;

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

const StyledItem = styled.div`
    width: 100%;
    height: 145px;
    border-bottom: 1px solid rgba(0,0,0,0.4);
    position: relative;
    padding: 20px 0px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: flex-start;
    h3 {
        margin: 0;
        font-size: 20px;
        font-weight: 600;
    }

`;