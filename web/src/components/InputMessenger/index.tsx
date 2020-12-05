import React from "react";
import styled from "styled-components";
import { ReactComponent as WhatsappIcon } from '../../assets/messengerIcons/whatsapp.svg';

const InputMessengerStyled = styled.div`
  .input-group{
    position: relative;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-wrap: wrap;
    flex-wrap: wrap;
    -ms-flex-align: stretch;
    align-items: stretch;
    width: 100%;
    border: 1px solid #ebedf8;
    box-sizing: border-box;

    button{
      background: white;
      border: 1px solid #ebedf8;
      display: inline-block;
      font-weight: 400;
      text-align: center;
      white-space: nowrap;
      vertical-align: middle;
      user-select: none;
      border: 1px solid transparent;
      padding: .6rem 1.2rem;
      font-size: 1rem;
      line-height: 1.5;
      border-radius: 3px;
      transition: color .15s ease-in-out,background-color .15s ease-in-out,border-color .15s ease-in-out,box-shadow .15s ease-in-out;
    }
    
    input{
      flex: 1 1 auto;
      border: inherit;
      padding: .6rem 1rem;
      font-family: inherit;
      font-size: 1rem;
    }
  }
  textarea{
    width: 100%;
    border: inherit;
    border: 1px solid #ebedf8;
    box-sizing: border-box;
    padding: .6rem 1rem;
    font-family: inherit;
    font-size: 1rem;
    resize: vertical;
  }
`

export interface InputMessengerProps {

}

export const InputMessenger: React.FC<InputMessengerProps> = ({

}) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><WhatsappIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="WhatsApp phone number with country code ( +1...)" />
      </div>
      <textarea placeholder="Predefined text: e.g Give me further information about..."/>
    </InputMessengerStyled>
  )
}