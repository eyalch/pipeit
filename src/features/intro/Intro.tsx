import { Button, Typography } from '@material-ui/core'
import React, { useState } from 'react'
import styled from 'styled-components/macro'
import CodeInput from './CodeInput'
import NewCode from './NewCode'

const StyledButtonsContainer = styled.div`
  display: flex;
  margin: ${(p) => p.theme.spacing(3, -1)};

  button {
    flex-grow: 1;
    flex-basis: 0;
    margin: ${(p) => p.theme.spacing(0, 1)};
  }
`

enum CodeSection {
  Enter,
  Create,
}

const Intro = () => {
  const [activeSection, setActiveSection] = useState<CodeSection>()

  const randomNumber = Math.floor(Math.random() * 9999)
  const randomCode = randomNumber.toString().padStart(4, '0')

  return (
    <>
      <Typography variant="h1">Got a code?</Typography>
      <Typography variant="subtitle2" gutterBottom>
        (from another device)
      </Typography>

      <StyledButtonsContainer>
        <Button
          variant="contained"
          size="large"
          color="primary"
          disableElevation
          disabled={activeSection === CodeSection.Enter}
          onClick={() => setActiveSection(CodeSection.Enter)}
        >
          Yes
        </Button>
        <Button
          variant="outlined"
          size="large"
          color="primary"
          disableElevation
          disabled={activeSection === CodeSection.Create}
          onClick={() => setActiveSection(CodeSection.Create)}
        >
          No
        </Button>
      </StyledButtonsContainer>

      {activeSection === CodeSection.Enter ? (
        <CodeInput />
      ) : activeSection === CodeSection.Create ? (
        <NewCode code={randomCode} />
      ) : null}
    </>
  )
}

export default Intro
