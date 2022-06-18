import Slider from '@material-ui/core/Slider';
import { styled } from '@mui/material/styles';

export const RangeSlider = ({value, setValue, min, max}) => {
    const handleChange = (event, newValue) => {
		setValue(newValue);
	};

    return ( <CustomSlider
                value={value}
                step={1}
                min={min}
                max={max}
                //className="floor_slid"
                onChange={handleChange}
                valueLabelDisplay="on"
            />
    )
}

const CustomSlider = styled(Slider)({
	color: '#424243',
	height: 8,
	//width: "350px",
	marginLeft: 20,
	'& .MuiSlider-track': {
		border: 'none',
		backgroundColor: '#0CA5E6',
	},
	'& .MuiSlider-rail': {
        backgroundColor: '#000000',
    },
	'& .MuiSlider-valueLabel span': {
		backgroundColor: "transparent !important",
		position: "absolute",
        color: "#fff",
		width: "28px",
		height: "28px",
		lineHeight: "28px",
		fontSize: "16px",
		fontWeight: 250,
		top: 26,
		left: 0,
		textAlign: "center",
	},
	'& .MuiSlider-thumb': {
		backgroundColor: "#025294",
		'&:before': {
			display: 'none',
		},
		width: "28px",
		height: "28px",
		textAlign: "center",
		fontSize: "16px",
		color: "#fff",
		fontWeight: 250,
		lineHeight: "56px",
		marginLeft: "-14px",
		marginTop: "-15px",
		cursor: "pointer",
		boxShadow: "0 0 0 10px transparent !important",
	},
	'& .MuiSlider-thumb:hover': {

		boxShadow: "0 0 0 10px transparent !important",
	}

});
