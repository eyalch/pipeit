import { Typography, TypographyProps } from '@material-ui/core'
import React from 'react'
import styled from 'styled-components/macro'

const _Brand: React.FC<TypographyProps> = (props) => (
  <Typography component="header" {...props} />
)
const StyledBrand = styled(_Brand)`
  font-size: 4.5rem;
  padding: 8vh 0;
`

const Header = () => <StyledBrand>PipeIt</StyledBrand>

export default Header
