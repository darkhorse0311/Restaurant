import React from 'react';
import { shallow } from 'enzyme';
import Navigation from '../../components/Navigation/index';

const props = {
    center: [-74.0060, 40.7128], 
    getLocations: jest.fn(), 
    loading: false, 
    permission: false,
    setCompact: jest.fn(),
    compact: false,
};

describe('Navigation Component', () => {
    it('renders without crashing', () => {
        const wrapper = shallow(<Navigation {...props}/>);
        expect(wrapper.exists()).toBe(true)
    });

    it('title exists', () => {
        const wrapper = shallow(<Navigation {...props}/>);
        expect(wrapper.find('h1').length).toEqual(1)
    });

    it('fa-redo-alt icon exists if permissions is true', () => {
        const wrapper = shallow(<Navigation {...props}/>);
        expect(wrapper.find('.fa-redo-alt').length).toEqual(1)
    });

    it('fa-spinner icon exists if loading is true', () => {
        const wrapper = shallow(<Navigation {...props} loading={true}/>);
        expect(wrapper.find('.fa-spinner').length).toEqual(1)
    });
    
    it('geo location warning shows with no props', () => {
        const wrapper = shallow(<Navigation/>);
        expect(wrapper.find('.geo-off').length).toEqual(1);
    });

    it('geo location warning shows when permissions is false', () => {
        const wrapper = shallow(<Navigation {...props} />);
        expect(wrapper.find('.geo-off').length).toEqual(1);
    });

    it('geo location warning does not show on permissions true', () => {
        const wrapper = shallow(<Navigation {...props} permission={true}/>);
        expect(wrapper.find('.geo-off').length).toEqual(0);
    });

    it('clicking refresh button calls getLocations', () => {
        const wrapper = shallow(<Navigation {...props}/>);
        wrapper.find('.fa-redo-alt').simulate('click');
        expect(props.getLocations).toHaveBeenCalledTimes(1);;
    });

    it('clicking compact button calls setCompact', () => {
        const wrapper = shallow(<Navigation {...props} compact={true}/>);
        wrapper.find('.fa-buffer').simulate('click');
        expect(props.setCompact).toHaveBeenCalledTimes(1);;
    });

});
