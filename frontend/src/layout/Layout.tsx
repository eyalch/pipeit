import { Container } from "@material-ui/core"
import React from "react"
import styled from "styled-components/macro"
import Header from "./Header"

const StyledLayout = styled(Container)`
  text-align: center;
`

const Layout: React.FC = ({ children }) => (
  <StyledLayout>
    <Header />

    <main>{children}</main>
  </StyledLayout>
)

export default Layout
