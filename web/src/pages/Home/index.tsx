import React from "react";
import styled from "styled-components";
import { ElementMessenger, ElementType, MessengerType, Sizes } from "../../components/ElementMessenger";
import { ElementBasic } from "../../components/ElementBasic";
import BasicLayout from "../../components/BasicLayout";
import { ElementWrapper } from "../../components/ElementWrapper";
import { InputMessenger, InputMessengerType } from "../../components/InputMessenger";

import { ReactComponent as CameraIcon } from '../../assets/camera.svg';
import { ReactComponent as PlusIcon } from '../../assets/plus.svg';

const ContainerStyled = styled.div`
  max-width: 386px;
  margin: 0 auto;
  padding-bottom: 16px;
  padding-top: 16px;
`;

const SvgStyled = styled.svg`
  position: relative;
  display: block;
  font-size: 2.7em;
  width: 1em;
  height: 100%;
  line-height: 100%;
  margin: 0 auto;
  margin-top: -0.1rem;
  color: #ffffff;
`;

const SvgContainerStyled = styled.div`
  border-color: rgba(255, 255, 255, 0.5);
  border: 2px dashed #ffffff;
  border-radius: 50px;
  width: 100px;
  height: 100px;
  margin: 0 auto;
  margin-bottom: 15px;
`;

const TitleStyled = styled.div`
  width: 100%;
  text-align: center;
  h1 {
    border-bottom: 2px dashed #fffdfd;
    display: inline-block;
    font-size: 28px;
    margin: 0;
    padding: 0;
    margin-bottom: 16px;
    text-align: center;
    color: white;
  }
`;

const LogoStyled = styled.div`
  text-align: center;
  font-family: Montserrat;
  font-weight: bolder;
  text-transform: uppercase;
  position: relative;
  line-height: 1em;
  color: #ffffff;
  -webkit-transition: all .62s;
  -o-transition: all .62s;
  transition: all .62s;
  padding: 16px 24px;
`

const MessengerSheetStyled = styled.div`
  transform: translateZ(0);
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  width: 100%;
  min-height: 100%;
  padding-top: 23vh;
  -webkit-overflow-scrolling: touch;
  overflow-x: hidden;
  overflow-y: scroll;
  background: rgba(0,0,0,.3);
  margin-bottom: -44px;
  z-index: 10;
  -webkit-transition: background-color .3s;
  -o-transition: background-color .3s;
  transition: background-color .3s;
  .table-align{
    color: #000;
    position: relative;
    bottom: 0;
    margin: 0 auto;
    padding: 2rem;
    background: #fff;
    border-radius: 24px;
    padding-bottom: calc(2em + 12px);
    margin-bottom: 20px;
    padding-bottom: calc(constant(safe-area-inset-bottom) + 2em + 12px);
    padding-bottom: calc(env(safe-area-inset-bottom) + 2em + 12px);
    -webkit-box-shadow: 0 -10px 25px rgba(0,0,0,.1);
    box-shadow: 0 -10px 25px rgba(0,0,0,.1);
    width: 100vw;
    max-width: 700px;
  }
`



const dataDummy = [
  {
    type: "messenger",
    messenger_type: "whatsapp",
    value: "081313131313",
    text: "text gua lho",
  },
  {
    type: "messenger",
    messenger_type: "facebook",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "telegram",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "skype",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "viber",
    value: "billysutomo",
  },
  {
    type: "messenger",
    messenger_type: "email",
    value: "billysutomo.53@gmail.com",
    subject: "subject email",
  },
  {
    type: "messenger",
    messenger_type: "phone",
    value: "081313131313",
  },
];

const Component: React.FC = () => {

  const findSize = (value: number, remainder: number): Sizes => {
    if ((remainder % 2 === 0) && value <= remainder) {
      return Sizes.medium
    } else if ((remainder % 1 === 0) && value <= remainder) {
      return Sizes.large
    } else {
      return Sizes.small
    }
  }

  const renderMessenger = () => {
    const remainder: number = (dataDummy.length % 3)
    return dataDummy.map((a, i) => {
      return <ElementMessenger
        key={i}
        active
        size={findSize(i + 1, remainder)}
        elementType={a.type as ElementType}
        messengerType={a.messenger_type as MessengerType} />;
    });
  };

  return (
    <BasicLayout backgroundColor="#cf8383">
      <ContainerStyled>
        <SvgContainerStyled>
          <SvgStyled>
            <CameraIcon />
          </SvgStyled>
        </SvgContainerStyled>
        <TitleStyled>
          <h1>Title here</h1>
        </TitleStyled>
        <ElementWrapper>
          {renderMessenger()}
        </ElementWrapper>
        <ElementBasic><PlusIcon width="0.75em" height="1em" /> Add Block</ElementBasic>
        <ElementBasic><PlusIcon width="0.75em" height="1em" /> Social Links</ElementBasic>
        <LogoStyled>Chocolate Waffle</LogoStyled>
        <MessengerSheetStyled>
          <div className="table-align">
            <div style={{ textAlign: "center" }}>Messenger</div>
            <InputMessenger messengerType={InputMessengerType.WHATSAPP} />
            <InputMessenger messengerType={InputMessengerType.FACEBOOK} />
            <InputMessenger messengerType={InputMessengerType.TELEGRAM} />
            <InputMessenger messengerType={InputMessengerType.SKYPE} />
            <InputMessenger messengerType={InputMessengerType.VIBER} />
            <InputMessenger messengerType={InputMessengerType.EMAIL} />
            <InputMessenger messengerType={InputMessengerType.PHONE} />
          </div>
        </MessengerSheetStyled>
      </ContainerStyled>
    </BasicLayout>
  );
};

export default Component;
