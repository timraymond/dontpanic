{
  description = "A flake for a PoC for globally capturing panics";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-24.11";
  };

  outputs = { self, nixpkgs }:
  let
    supportedSystems = [ "x86_64-linux" "x86_64-darwin" ];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
  in
  {
    packages = forAllSystems (system:
      let
        pkgs = nixpkgsFor.${system};
      in
      {
        default = pkgs.buildGoModule {
          pname = "dontpanic";
          version = "0.0.1";
          src = ./.;
          vendorHash = null;
          proxyVendor = true;
        };
    });
    devShells = forAllSystems (system:
      let
        pkgs = nixpkgsFor.${system};
      in
      {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
          ];
        };
      }
    );
  };
}
