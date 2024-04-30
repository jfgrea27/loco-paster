{
  description = "A flake for loco-paster.";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        p = import nixpkgs {
          inherit system;
        };
      in
        {
          devShells = rec {
            default = nixpkgs.legacyPackages.${system}.mkShell {
              packages = [
                # golang specific
                p.go
                p.gopls
                p.go-tools
                p.errcheck
                p.gofumpt
                # react specific
                p.nodejs_21
                p.nodePackages.prettier
              ];
              shellHook=''
                export GOROOT=${p.go.outPath}/share/go
                unset GOPATH
          '';
            };
          };
        }
    );
}
