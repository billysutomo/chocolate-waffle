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
}

export const InputMessenger: React.FC<InputMessengerProps> = ({
  messengerType
}) => {

  const findIcon = (type: InputMessengerType): React.SVGProps<SVGAElement> => {
    if (type === InputMessengerType.FACEBOOK) {
      return <FacebookIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} />
    } else if (type === InputMessengerType.TELEGRAM) {
      return <TelegramIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} />
    } else if (type === InputMessengerType.SKYPE) {
      return <SkypeIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} />
    } else if (type === InputMessengerType.VIBER) {
      return <ViberIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} />
    } else {
      return <PhoneIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} />
    }
  }

  const renderInputMessenger = () => {
    if (messengerType == InputMessengerType.WHATSAPP) {
      return <InputWA />
    } else if (messengerType == InputMessengerType.EMAIL) {
      return <InputEmail />
    } else {
      const icon = findIcon(messengerType);

      return <InputStandard
        icon={icon}
      />
    }
  }

  return (
    <div style={{ display: "flex", marginBottom: "1rem" }}>
      <BarIcon height="17px" width="19px" style={{ alignSelf: "center", padding: ".4em .8em" }} />
      {renderInputMessenger()}
    </div>
  )
}

const InputWA: React.FC = ({

}) => {
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

const InputEmail: React.FC = ({

}) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button><EmailIcon fill="red" height="24px" width="24px" style={{ verticalAlign: "middle" }} /></button>
        <input placeholder="Email address" />
      </div>
      <textarea placeholder="Predefined text: e.g Give me further information about..." />
    </InputMessengerStyled>
  )
}

interface InputStandardProps {
  icon: React.SVGProps<SVGAElement>
}

const InputStandard: React.FC<InputStandardProps> = ({
  icon
}) => {
  return (
    <InputMessengerStyled>
      <div className="input-group">
        <button>{icon}</button>
        <input placeholder="WhatsApp phone number with country code ( +1...)" />
      </div>
    </InputMessengerStyled>
  )
}