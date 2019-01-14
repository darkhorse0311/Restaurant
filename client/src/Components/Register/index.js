import React, { Component } from 'react';
import styled from 'styled-components';
import axios from 'axios';


class Register extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: ''
        }
    }

    // componentDidUpdate() {
    //     if (this.props.authenticated) {
    //         this.props.history.push('/app');
    //     }
    // }

    changeHandler = e => {
        e.preventDefault();
        const { name, value } = e.target;
        this.setState({[name]: value})
    }

    responseGoogle = (response) => {
        // console.log(response);
        // if (response.Zi && response.Zi.id_token) {
        //     const config = {
        //         headers: {
        //           authorization: response.Zi.id_token,
        //         }
        //     }
        //     axios.get('http://localhost:9001/auth/oauth/google', config);
        // }
    }

    render() {
        return (
            <Container>
                <Wrapper
                    onSubmit={(e) => {
                        e.preventDefault();
                        this.props.authRegister({...this.state})
                    }}>
                    <StyledLabel htmlFor="username">Username</StyledLabel>
                    <StyledInput 
                        type="text" 
                        name="username" 
                        id="username"
                        value={this.state.username}
                        onChange={this.changeHandler}
                        placeholder="Username"
                    />

                    <StyledLabel htmlFor="password">Password</StyledLabel>
                    <StyledInput 
                        type="text" 
                        name="password" 
                        id="password"
                        value={this.state.password}
                        onChange={this.changeHandler}
                        placeholder="Password"
                    />
                    
                    <StyledButton type="submit">Submit</StyledButton>
                    {/* {
                        this.props.error !== null
                        ? <ErrorWrapper>Error Registering</ErrorWrapper>
                        : null
                    } */}
                </Wrapper>
            </Container>
        );
    }
}


export default Register;


const StyledButton = styled.button`
    margin-top: 20px;
    height: 35px;
    border-radius: 6px;
    line-height: 35px;
    font-size: 16px;
    cursor: pointer;
    width: 100%;
`;

const StyledInput = styled.input`
    height: 35px;
    border: 1px solid rgba(0,0,0,0.2);
    border-radius: 6px;
    outline: none;
    line-height: 35px;
    padding: 0px 10px;
    font-size: 16px;
    width: 100%;
`;

const StyledLabel = styled.label`
    line-height: 35px;
    width: 100%;
`;

const Container = styled.div`
    width: 100vw;
    height: 100vh;
    position: fixed;
    left: 0;
    top: 0;
    z-index: 3;
    background-color: rgba(0,0,0,0.4);
    display: flex;
    justify-content: center;
    align-items: center;
`;

const Wrapper = styled.form`
    width: 300px;
    background-color: white;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 30px 25px 40px;
    box-shadow: 0 2px 2px 0 rgba(0,0,0,0.14), 0 3px 1px -2px rgba(0,0,0,0.12), 0 1px 5px 0 rgba(0,0,0,0.2);
`

const ErrorWrapper = styled.span`
    color: rgba(255,0,0,0.8);
    font-size: 16px;
    font-weight: 600;
    margin: 10px 0;
`;