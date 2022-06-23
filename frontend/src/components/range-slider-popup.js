import Slider from '@material-ui/core/Slider';
import { styled } from '@mui/material/styles';

export const RangeSliderPopup = ({step, min, max, value, postfix, callback}) => {
	const num_word = (value, words) => {  
		value = Math.abs(value) % 100; 
		var num = value % 10;
		if(value > 10 && value < 20) return words[2]; 
		if(num > 1 && num < 5) return words[1];
		if(num == 1) return words[0]; 
		return words[2];
	}
    return ( <CustomSlider
                step={step}
                min={min}
                max={max}
                //className="floor_slid"
                onChange={callback}
                valueLabelDisplay="on"
				valueLabelFormat={postfix.length>1?`${value} ${num_word(value, postfix)}`:`${value}${postfix[0]}`}
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
