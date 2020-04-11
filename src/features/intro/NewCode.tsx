import {
  CircularProgress,
  Typography,
  TypographyProps,
} from '@material-ui/core'
import React from 'react'
import styled from 'styled-components/macro'

const _CodeDigit: React.FC<TypographyProps> = (props) => (
  <Typography component="span" {...props} />
)
const StyledCodeDigit = styled(_CodeDigit)`
  font-size: 4.5rem;
  margin: ${(p) => p.theme.spacing(0, 1)};
`
const StyledCircularProgress = styled(CircularProgress)`
  margin: ${(p) => p.theme.spacing(2)}px;
`

type NewCodeProps = { code: string }

const NewCode: React.FC<NewCodeProps> = ({ code }) => (
  <div>
    {code.split('').map((digit, i) => (
      <StyledCodeDigit key={i}>{digit}</StyledCodeDigit>
    ))}
    <Typography>Enter this code on another device...</Typography>
    <StyledCircularProgress disableShrink />
  </div>
)

export default NewCode
