export const UPDATE_SELECTED = 'update_selected';


export const updateSelected = (place) => {
    return {
        type: UPDATE_SELECTED,
        payload: place
    }
}