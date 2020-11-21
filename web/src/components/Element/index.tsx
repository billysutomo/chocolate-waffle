import React from "react";
import styled, { css } from "styled-components";

import { ReactComponent as WaLogo } from '../../assets/whatsapp.svg';

const ElementStyled = styled.button<ElementProps>`
  border: 2px dashed #607d8b;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 370px;
  text-align: center;
  margin-bottom: 16px;
  background-color: ${(p) => p.backgroundColor};
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
    if (messengerType == MessengerType.WHATSAPP) {
      return (
        <span>
          <WaLogo height="24px" width="24px" style={{ verticalAlign: "middle" }} />
          <MessengerStyled>Whatsapps</MessengerStyled>
        </span>
      )
    } else if (messengerType == MessengerType.FACEBOOK) {
      return <WaLogo height="24px" width="24px" />
    } else if (messengerType == MessengerType.TELEGRAM) {
      return <WaLogo height="24px" width="24px" />
    } else if (messengerType == MessengerType.SKYPE) {
      return <WaLogo height="24px" width="24px" />
    } else if (messengerType == MessengerType.VIBER) {
      return <WaLogo height="24px" width="24px" />
    } else if (messengerType == MessengerType.EMAIL) {
      return <WaLogo height="24px" width="24px" />
    } else if (messengerType == MessengerType.PHONE) {
      return <WaLogo height="24px" width="24px" />
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
