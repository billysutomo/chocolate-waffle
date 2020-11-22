import React from "react";
import styled, { css } from "styled-components";

import { ReactComponent as WhatsappIcon } from '../../assets/messengerIcons/whatsapp.svg';
import { ReactComponent as FacebookIcon } from '../../assets/messengerIcons/facebook.svg';
import { ReactComponent as TelegramIcon } from '../../assets/messengerIcons/telegram.svg';
import { ReactComponent as SkypeIcon } from '../../assets/messengerIcons/skype.svg';
import { ReactComponent as ViberIcon } from '../../assets/messengerIcons/viber.svg';
import { ReactComponent as EmailIcon } from '../../assets/messengerIcons/email.svg';
import { ReactComponent as PhoneIcon } from '../../assets/messengerIcons/phone.svg';

const ElementStyled = styled.button<ElementProps>`
  border: 2px dashed #607d8b;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
  margin-bottom: 16px;
  background-color: ${(p) => p.backgroundColor};
  margin: 0 auto;
  margin-left: 8px;
  margin-right: 8px;
  margin-bottom: 16px;
  :hover {
    cursor: pointer;
  }
  :active {
    background: #2ab8b96b;
  }
  :focus {
    outline: 0;
  }
  ${(p) =>
    p.size === Sizes.medium &&
    css`
      width: 177px;
    `}
  ${(p) =>
    p.size === Sizes.small &&
    css`
      width: 112.66px;
    `}
  ${(p) =>
    p.active &&
    css`
      border: none;
    `}
`;

const MessengerStyled = styled.span`
  vertical-align: middle; 
  font-size: 16px;
  font-weight: bold;
  margin-left: .6em;
`

export enum Sizes {
  small,
  medium,
  large,
}

export enum ElementType {
  MESSENGER = "messenger",
  BLOCK = "block",
  SOCIAL_LINK = "social_link",
  NOTHING = "nothing",
}

export enum MessengerType {
  WHATSAPP = "whatsapp",
  FACEBOOK = "facebook",
  TELEGRAM = "telegram",
  SKYPE = "skype",
  VIBER = "viber",
  EMAIL = "email",
  PHONE = "phone"
}

export interface ElementProps {
  active: boolean;
  backgroundColor?: string;
  size?: Sizes;
  elementType: ElementType;
  messengerType: MessengerType;
}

export const Element: React.FC<ElementProps> = ({
  active,
  backgroundColor,
  size,
  elementType,
  messengerType
}) => {

  const renderMessengerIcon = () => {
    if (messengerType === MessengerType.WHATSAPP) {
      return (
        <span>
          <WhatsappIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Whatsapps</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.FACEBOOK) {
      return (
        <span>
          <FacebookIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Messenger</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.TELEGRAM) {
      return (
        <span>
          <TelegramIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Telegram</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.SKYPE) {
      return (
        <span>
          <SkypeIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Skype</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.VIBER) {
      return (
        <span>
          <ViberIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Viber</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.EMAIL) {
      return (
        <span>
          <EmailIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Email</MessengerStyled> : null}
        </span>
      )
    } else if (messengerType === MessengerType.PHONE) {
      return (
        <span>
          <PhoneIcon height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          {size === Sizes.large ? <MessengerStyled>Phone</MessengerStyled> : null}
        </span>
      )
    } else {
      return null
    }
  }

  return (
    <ElementStyled
      size={size}
      active={active}
      backgroundColor={backgroundColor}
      elementType={elementType}
      messengerType={messengerType}
    >
      {renderMessengerIcon()}
    </ElementStyled>
  );
};
