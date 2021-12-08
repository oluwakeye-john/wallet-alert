(function () {
  // scoped

  const Ethereum = window.ethereum;
  let address_text = "";
  let email_text = "";
  let currentForm = 0;

  const form1ButtonEL = document.querySelector("#form1-button");
  const form2ButtonEl = document.querySelector("#form2-button");
  const addressInputEl = document.querySelector("#address-input");

  const emailInputEl = document.querySelector("#email-input");
  const metamaskButtonEl = document.querySelector("#metamask-login-btn");

  const form1El = document.querySelector(".form-1");
  const form2El = document.querySelector(".form-2");

  const onMetamaskClick = async () => {
    if (Ethereum && Ethereum.isMetaMask) {
      const { error, account } = await getMetamaskAddress();

      if (error) {
        console.log(error);
      } else {
        addressInputEl.value = account;
        updateAddress(account);
      }
    } else {
      alert("Install metamask");
    }
  };

  const handleContinue = () => {
    if (currentForm === 0) {
      form1El.classList.remove("form--show");
      form2El.classList.add("form--show");

      currentForm += 1;
    } else {
      console.log({ email_text, address_text });
    }
  };

  const watchInput = () => {
    addressInputEl.addEventListener("input", (e) => {
      updateAddress(e.target.value);
    });

    emailInputEl.addEventListener("input", (e) => {
      console.log(";;");
      email_text = e.target.value;
      if (e.target.value) {
        setButton2(true);
      } else {
        setButton2(false);
      }
    });
  };

  const updateAddress = (input) => {
    address_text = input;
    setButton1(input);
  };

  const setButton1 = (enabled) => {
    form1ButtonEL.disabled = !enabled;
  };

  const setButton2 = (enabled) => {
    form2ButtonEl.disabled = !enabled;
  };

  const getMetamaskAddress = async () => {
    const request = {
      method: "eth_requestAccounts",
      params: [],
    };

    const defaultErrorMessage = "Error occurred";

    try {
      if (Ethereum.isConnected() && Ethereum.selectedAddress) {
        const account = Ethereum.selectedAddress;
        return { error: "", account };
      }

      const accounts = await Ethereum.request(request);

      if (accounts.length >= 1) {
        const account = accounts[0];
        return { error: "", account };
      } else {
        return defaultErrorMessage, "";
      }
    } catch (err) {
      return { error: err?.message || defaultErrorMessage, account: null };
    }
  };

  watchInput();

  metamaskButtonEl.addEventListener("click", onMetamaskClick);

  form1El.addEventListener("submit", (e) => {
    e.preventDefault();
    handleContinue();
  });

  form2El.addEventListener("submit", (e) => {
    e.preventDefault();
    handleContinue();
  });
})();
