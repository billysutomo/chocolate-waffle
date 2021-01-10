import React from "react";
import styled from "styled-components";
import { ReactComponent as BarIcon } from '../../assets/bar.svg';
import { ReactComponent as WhatsappIcon } from '../../assets/messengerIcons/whatsapp.svg';
import { ReactComponent as FacebookIcon } from '../../assets/messengerIcons/facebook.svg';
import { ReactComponent as TelegramIcon } from '../../assets/messengerIcons/telegram.svg';
import { ReactComponent as SkypeIcon } from '../../assets/messengerIcons/skype.svg';
import { ReactComponent as ViberIcon } from '../../assets/messengerIcons/viber.svg';
import { ReactComponent as EmailIcon } from '../../assets/messengerIcons/email.svg';
import { ReactComponent as PhoneIcon } from '../../assets/messengerIcons/phone.svg';

const InputMessengerStyled = styled.div`
  flex: auto;
  .input-group{
    position: relative;
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

export enum InputMessengerType {
  WHATSAPP = "whatsapp",
  FACEBOOK = "facebook",
  TELEGRAM = "telegram",
  SKYPE = "skype",
  VIBER = "viber",
  EMAIL = "email",
  PHONE = "phone"
}

export interface InputMessengerProps {
  messengerType: InputMessengerType;
  propRef: React.RefObject<HTMLDivElement>;
}

export const InputMessenger: React.FC<InputMessengerProps> = ({
  messengerType,
  propRef,
  ...props
}) => {

  const renderInputMessenger = (type: InputMessengerType) => {
    switch (type) {
      case InputMessengerType.WHATSAPP:
        return <InputWhatsapp />
      case InputMessengerType.FACEBOOK:
        return <InputFacebookMesseger />
      case InputMessengerType.TELEGRAM:
        return <InputTelegram />
      case InputMessengerType.SKYPE:
        return <InputSkype />
      case InputMessengerType.VIBER:
        return <InputViber />
      case InputMessengerType.EMAIL:
        return <InputEmail />
      case InputMessengerType.PHONE:
        return <InputPhone />
    }
  }

  return (
    <div style={{ display: "flex", marginBottom: "1rem" }} ref={propRef} {...props}>
      <BarIcon height="17px" width="19px" style={{ alignSelf: "center", padding: ".4em .8em" }} />
      {renderInputMessenger(messengerType)}
    </div>
  )
}

const InputWhatsapp: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><WhatsappIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="WhatsApp phone number with country code ( +1...)" />
      </div>
      <textarea placeholder="Predefined text: e.g Give me further information about..." />
    </InputMessengerStyled>
  )
}
const InputFacebookMesseger: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><FacebookIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Messenger/Facebook username" />
      </div>
    </InputMessengerStyled>
  )
}

const InputTelegram: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><TelegramIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Telegram username" />
      </div>
    </InputMessengerStyled>
  )
}

const InputSkype: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><SkypeIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Skype username" />
      </div>
    </InputMessengerStyled>
  )
}

const InputViber: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><ViberIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Viber number" />
      </div>
    </InputMessengerStyled>
  )
}

const InputEmail: React.FC = ({
}) => {
  return (
    <InputMessengerStyled >
      <div className="input-group">
        <button><EmailIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Email address" />
      </div>
      <textarea placeholder="Subject:" />
    </InputMessengerStyled>
  )
}

const InputPhone: React.FC = ({ }) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><PhoneIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Phone number with country code (+1...)" />
      </div>
    </InputMessengerStyled>
  )
}