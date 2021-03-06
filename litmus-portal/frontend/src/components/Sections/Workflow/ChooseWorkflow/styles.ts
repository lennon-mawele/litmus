import {
  makeStyles,
  TextField,
  withStyles,
  Theme,
  Button,
} from '@material-ui/core';
import { purple } from '@material-ui/core/colors';

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: 'rgba(255, 255, 255, 0.6)',
    height: '100%',
    width: '80%',
    margin: '0 auto',
    border: 1,
    borderColor: 'rgba(0, 0, 0, 0.05)',
    borderRadius: 3,
    flexDirection: 'column',
    paddingBottom: theme.spacing(5),
    paddingLeft: theme.spacing(3.125),
    paddingRight: theme.spacing(3.125),
  },

  cards: {
    paddingLeft: theme.spacing(1.875),
    paddingRight: theme.spacing(1.875),
  },

  heading: {
    marginTop: theme.spacing(3),
    fontFamily: 'Ubuntu',
    fontSize: '1.5rem',
    marginLeft: theme.spacing(2),
  },

  modalHeading: {
    marginTop: theme.spacing(1.5),
    fontFamily: 'Ubuntu',
    fontSize: '2.25rem',
  },

  description: {
    width: '90%',
    marginTop: theme.spacing(3.25),
    fontFamily: 'Ubuntu',
    fontSize: '1rem',
    marginLeft: theme.spacing(2),
  },

  horizontalLine: {
    marginTop: theme.spacing(7.5),
    marginBottom: theme.spacing(4.5),
  },

  paddedTop: {
    marginTop: theme.spacing(5),
  },

  inputArea: {
    marginTop: theme.spacing(3),
    display: 'flex',
    textDecoration: 'none',
    flexDirection: 'column',
    paddingLeft: theme.spacing(2),
    paddingBottom: theme.spacing(2.2),
    borderRadius: 3,
    color: 'rgba(0, 0, 0, 0.2)',
    border: '1px solid rgba(0, 0, 0, 0.2)',
  },

  inputAreaDescription: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    display: 'flex',
    textDecoration: 'none',
    flexDirection: 'column',
    paddingLeft: theme.spacing(2),
    paddingBottom: theme.spacing(2.2),
    borderRadius: 3,
    color: 'rgba(0, 0, 0, 0.2)',
    border: '1px solid rgba(0, 0, 0, 0.2)',
  },

  textfieldworkflowname: {
    marginTop: theme.spacing(1),
    paddingLeft: theme.spacing(2),
  },

  totalWorkflows: {
    fontSize: '1rem',
    fontWeight: 500,
  },

  textfieldworkflowdescription: {
    marginTop: theme.spacing(2),
    paddingLeft: theme.spacing(2),
  },

  inputDiv: {
    marginTop: theme.spacing(5),
    marginRight: theme.spacing(4.5),
  },

  saved: {
    width: '25rem',
    marginTop: theme.spacing(6),
    fontFamily: 'Ubuntu',
    lineHeight: '130%',
    fontWeight: 500,
    fontSize: '1rem',
    color: theme.palette.primary.dark,
  },

  testType: {
    fontSize: '1rem',
    fontFamily: 'Ubuntu',
    paddingRight: theme.spacing(1.25),
  },

  resize: {
    fontSize: '1rem',
  },

  resizeName: {
    fontSize: '0.875rem',
  },

  modalContainerName: {
    paddingBottom: theme.spacing(0.625),
  },

  modalContainerBody: {
    position: 'relative',
    height: '70%',
    paddingLeft: theme.spacing(18),
    paddingRight: theme.spacing(14),
    paddingBottom: theme.spacing(0.625),
  },

  buttons: {
    display: 'flex',
    justifyContent: 'space-between',
  },

  colorButton: {
    textTransform: 'none',
    fontSize: '0.75rem',
    lineHeight: 2.5,
    borderRadius: 3,
  },

  selectionName: {
    fontFamily: 'Ubuntu',
    fontSize: '1rem',
  },

  buttonCancel: {
    textTransform: 'none',
    fontSize: '0.75rem',
    lineHeight: 2.5,
    borderRadius: 3,
    color: 'black',
    border: '1px solid #5B44BA',
  },

  buttonSave: {
    textTransform: 'none',
    fontSize: '0.75rem',
    lineHeight: 2.5,
    borderRadius: 3,
    width: '6.6rem',
    marginRight: theme.spacing(4.375),
  },
}));

export const CssTextField = withStyles({
  root: {
    '& label.MuiInputLabel-root': {
      color: 'rgba(0, 0, 0, 0.6)',
      fontSize: 16,
    },
  },
})(TextField);

export const ColorButton = withStyles((theme: Theme) => ({
  root: {
    color: theme.palette.getContrastText(purple[500]),
    backgroundColor: theme.palette.primary.dark,
  },
  '&:hover': {
    color: theme.palette.getContrastText(purple[500]),
    backgroundColor: theme.palette.primary.dark,
  },
}))(Button);

export default useStyles;
