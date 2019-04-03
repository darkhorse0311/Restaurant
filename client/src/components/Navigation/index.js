import React from 'react';
import styled, { keyframes, css } from 'styled-components'

const Navigation = ({center, getLocations, loading, permission, compact, setCompact}) => {
    return (
        <StyledHeader compact={compact}>
            <h1>Carbtographer</h1>
            <span className="header-actions">
                {
                    loading
                        ? (<i className="fas fa-spinner spinner" />)
                        : (<i
                            className="fas fa-redo-alt"
                            onClick={() => {
                                getLocations(center[0], center[1])
                            }}
                        />)
                }
                <i 
                    className="fab fa-buffer"
                    onClick={() => setCompact(!compact)}
                ></i>
            </span>
            {
                !permission ? (
                    <span className="geo-off">
                        <i className="fas fa-map-marker-alt"/>
                        Location Permissions: OFF
                    </span>
                ) : null
            }
        </StyledHeader>
    );
}

export default Navigation;

const rotation = keyframes`
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(359deg);
    }
`

const StyledHeader = styled.div`
    width: 100%;
    @media (min-width: 650px) {
        width: calc(100vw - 325px)
    };
    height: 60px;
    position: absolute;
    z-index: 2;
    top: 0;
    left: 0;
    display: flex;
    padding: 0 20px;
    justify-content: space-between;
    align-items: center;
    color: rgba(255,255,255,0.74);
    .header-actions {
        display: flex;
        /* flex-direction: column; */
        /* margin-top: 78px; */
        margin-top: 18px;
        justify-content: flex-start;
        align-items: center;
        i {
            color: rgba(255,255,255,0.70);
            font-size: 26px;
            height: 60px;
            width: 60px;
            line-height: 60px;
            text-align: center;
            cursor: pointer;
            transition: font-size 200ms ease-in-out;
            :hover {
                font-size: 28px;
            }
            &.fa-buffer {
                ${({compact}) => compact === true && css`
                    color: white;
                `}
            }
            &.spinner {
                animation: ${rotation} 2.5s infinite linear;
            }
        }
    }
    h1 {
        line-height: 60px;
        margin: 0;
    }
    .geo-off {
        font-size: 16px;
        position: absolute;
        left: 20px;
        bottom: -15px;
        i {
            font-size: 16px;
            height: 20px;
            width: 30px;
        }
    }
`;

