import { TextField } from "@material-ui/core"
import React, { useRef, useState } from "react"
import styled from "styled-components/macro"

const CODE_LENGTH = 4
const INITIAL_CODE = Array<string>(CODE_LENGTH).fill("")

const StyledTextField = styled(TextField)`
  width: 56px;
  height: auto;
  margin: ${(p) => p.theme.spacing(0, 1)};

  input {
    font-size: 4.5rem;
    text-align: center;
  }
`

const CodeInput = () => {
  const [code, setCode] = useState(INITIAL_CODE)

  const inputsRef = useRef<HTMLInputElement[]>([])

  const onComplete = (code: string) => {
    // Submit the entered code
    alert(code)

    // Reset the code & clear the input
    setCode(INITIAL_CODE)
  }

  const handleChange = (value: string, index: number) =>
    setCode((code) => {
      const newCode = [...code.slice(0, index), value, ...code.slice(index + 1)]

      const codeStr = newCode.join("")

      // If the code is fully entered, then call the `onComplete` callback;
      // otherwise, if the changed input isn't empty and isn't the last one, then focus
      // the next one
      if (codeStr.length === CODE_LENGTH) {
        onComplete(codeStr)
      } else if (value && index < CODE_LENGTH - 1) {
        inputsRef.current[index + 1].focus()
      }

      return newCode
    })

  const handleBackspace = (index: number) => {
    // Do nothing for the first input and for non-empty inputs
    if (index === 0 || code[index] !== "") return

    // Clear the previous digit/input
    setCode((code) => [...code.slice(0, index - 1), "", ...code.slice(index)])

    // Focus the previous input
    inputsRef.current[index - 1].focus()
  }

  return (
    <>
      {code.map((digit, i) => (
        <StyledTextField
          key={i}
          variant="standard"
          type="tel"
          inputProps={{ maxLength: 1 }}
          inputRef={(el) => (inputsRef.current[i] = el)}
          value={digit}
          onChange={(event) => handleChange(event.target.value, i)}
          onKeyDown={(event) => event.key === "Backspace" && handleBackspace(i)}
          onFocus={(event) => event.target.select()}
          autoFocus={i === 0}
        />
      ))}
    </>
  )
}

export default CodeInput
