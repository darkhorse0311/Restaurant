import React from 'react';
import { shallow, mount } from 'enzyme';

import App from '../components/App';

// import rootReducer from '../reducers';
// import { Provider } from 'react-redux'
// import { createStore, applyMiddleware } from 'redux';
// import { MemoryRouter } from 'react-router';
// import thunk from 'redux-thunk';
// import Map from '../components/Map/index'
// import Items from '../components/Items/index'
// let store = createStore(rootReducer, applyMiddleware(thunk));

describe('App component', () => {
    it('renders without crashing', () => {
        const wrapper = shallow(<App/>);
        expect(wrapper.exists()).toBe(true)
    });
    // it('renders Map and Items component', () => {
    //     const wrapper = mount(
    //         <Provider store={store}>
    //             <MemoryRouter initialEntries={[ '/' ]}>
    //                 <App/>
    //             </MemoryRouter>
    //         </Provider>
    //     );
    //     expect(wrapper.find(Map)).toHaveLength(0);
    //     expect(wrapper.find(Items)).toHaveLength(0);
    // });
});
