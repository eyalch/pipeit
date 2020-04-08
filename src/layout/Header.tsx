import React from 'react'
import styled from 'styled-components'

const StyledHeader = styled.header`
  padding: 20% 0;
`
const StyledLogo = styled.h1`
  font-family: Inconsolata, monospace;
  font-size: 4.5rem;
  font-weight: normal;
  text-align: center;
  margin: 0;
`

const Header = () => (
  <StyledHeader>
    <StyledLogo>PipeIO</StyledLogo>
  </StyledHeader>
)

export default Header
